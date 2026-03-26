import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { MemoryRouter, useNavigate } from 'react-router-dom'
import Login from './Login'

const mockNavigate = vi.fn()

vi.mock('react-router-dom', async () => {
  const actual = await vi.importActual('react-router-dom')
  return {
    ...actual,
    useNavigate: () => mockNavigate,
  }
})

vi.mock('../components/ParabolicBackground', () => {
  return {
    default: () => null,
  }
})

function renderLogin() {
  return render(
    <MemoryRouter>
      <Login />
    </MemoryRouter>
  )
}

beforeEach(() => {
  vi.resetAllMocks()
  global.fetch = vi.fn()
  localStorage.clear()
})

describe('Login', () => {

  it('renders username and password fields', () => {
    renderLogin()
    expect(screen.getByPlaceholderText('Username')).toBeInTheDocument()
    expect(screen.getByPlaceholderText('Password')).toBeInTheDocument()
  })

  it('stores token and navigates on successful login', async () => {
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ token: 'abc123', username: 'admin' }),
    })

    renderLogin()
    await userEvent.type(screen.getByPlaceholderText('Username'), 'admin')
    await userEvent.type(screen.getByPlaceholderText('Password'), 'secret{Enter}')

    await waitFor(() => {
      expect(localStorage.getItem('token')).toBe('abc123')
      expect(localStorage.getItem('username')).toBe('admin')
    })

    await waitFor(() => {
      expect(mockNavigate).toHaveBeenCalledWith('/dashboard')
    }, { timeout: 3000 })
  })

  it('displays error message on failed login', async () => {
    global.fetch = vi.fn().mockResolvedValue({
      ok: false,
      json: async () => ({ error: 'Invalid credentials' }),
    })

    renderLogin()
    await userEvent.type(screen.getByPlaceholderText('Username'), 'admin')
    await userEvent.type(screen.getByPlaceholderText('Password'), 'wrong{Enter}')

    await waitFor(() => {
      expect(screen.getByText('Invalid credentials')).toBeInTheDocument()
    })
  })

  it('displays network error when fetch fails', async () => {
    global.fetch = vi.fn().mockRejectedValue(new Error('Network error'))

    renderLogin()
    await userEvent.type(screen.getByPlaceholderText('Username'), 'admin')
    await userEvent.type(screen.getByPlaceholderText('Password'), 'secret{Enter}')

    await waitFor(() => {
      expect(screen.getByText('Unable to connect to the server.')).toBeInTheDocument()
    })
  })

  it('dicplays field validation errors', async () => {
    global.fetch = vi.fn().mockResolvedValue({
      ok: false,
      json: async () => ({
        error: 'Validation failed',
        fields: { username: 'too short', password: 'required' },
      }),
    })

    renderLogin()
    await userEvent.type(screen.getByPlaceholderText('Username'), 'a')
    await userEvent.type(screen.getByPlaceholderText('Password'), '{Enter}')

    await waitFor(() => {
      expect(screen.getByText(/Validation failed/)).toBeInTheDocument()
      expect(screen.getByText(/username - too short/)).toBeInTheDocument()
      expect(screen.getByText(/password - required/)).toBeInTheDocument()
    })
  })
})
