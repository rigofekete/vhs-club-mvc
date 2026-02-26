import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import './Login.css'
import Header from '../components/Header'
import InputField from '../components/InputField'

function Login() {
  const navigate = useNavigate()
  const [logged, setLogin] = useState(false);
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');


  const inputFields = [
    {
      className: "Username", type: "text", label: "Username",
      onChange: (e) => setUsername(e.target.value)
    },
    {
      className: "Password", type: "password", label: "Password",
      onChange: (e) => setPassword(e.target.value)
    },
  ];

  const handleLogin = async (event) => {
    event.preventDefault();
    setError('');

    try {
      const response = await fetch('http://localhost:8080/users/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ "username": username, "password": password })
      });

      if (!response.ok) {
        const errData = await response.json();
        let errorMsg = errData.error || 'Login failed. Please try again.';
        if (errData.fields) {
          errorMsg += ': ' + Object.entries(errData.fields)
            .map(([field, msg]) => `${field} - ${msg}`)
            .join(', ');
        }
        setError(errorMsg)
        return;
      }

      const data = await response.json();
      localStorage.setItem('token', data.token);
      localStorage.setItem('username', data.username);
      setLogin(true);
      setTimeout(() => {
        navigate('/dashboard');
      }, 3000);
    } catch (err) {
      setError('Unable to connect to server.');
    }
  };


  return (
    <>
      <div className="Login">
        < Header logged={logged} />
        {logged ? (
          <p className="welcome-text"> Welcome, {username}</p>
        ) : (
          <form onSubmit={handleLogin} >
            {inputFields.map((inputField, index) => (
              <InputField
                key={index}
                className={inputField.className}
                type={inputField.type}
                label={inputField.label}
                onChange={inputField.onChange}
              />
            ))}
            <button className="button" type="submit">Login</button>
            <p className="error-message">
              {error || '\u00A0'}
            </p>
          </form>
        )}
      </div >
    </>
  )
}

export default Login

