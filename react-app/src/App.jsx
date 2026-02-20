import { useState, useEffect } from 'react'
import './App.css'
import vhsLogo from './assets/vhsLogo.svg'
// import Header from './components/Header'
import InputField from './components/InputField'

function App() {
  const [logged, setLogin] = useState(false);
  const [users, setUsers] = useState([]);
  const [username, setUsername] = useState('');
  const [error, setError] = useState('');


  const inputFields = [
    { className: "Username", type: "text", label: "Username" },
    { className: "Password", type: "text", label: "Password" },
  ];

  const handleChange = (event) => {
    setUsername(event.target.value);
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080/users');
        const jsonData = await response.json();
        setUsers(jsonData);
      } catch (err) { console.error('Failed to fetch users:', err); }
    };
    fetchData();
  }, []);


  const handleLogin = (event) => {
    event.preventDefault();

    const userExists = users.some((user) => user.username === username);

    console.log('Users fetched', users);

    if (userExists) {
      setLogin(true);
      console.log("Login complete!")
    } else {
      setError('Username not found. Please try again.');
    }
  };

  return (
    <>
      <div className="App">
        < img src={vhsLogo}
          className={logged ? "logo-blink" : "logo}"}
          alt="Vhs Logo"
        />
        {logged ? (
          <p className="welcome-text"> Welcome, {username}</p>
        ) : (
          <form onSubmit={handleLogin} onChange={handleChange}>
            {inputFields.map((inputField, index) => (
              <InputField
                key={index}
                className={inputField.className}
                type={inputField.type}
                label={inputField.label}
              />
            ))}
            <button className="button" type="submit">Login</button>
            {error && <p className="error-message">{error}</p>}
          </form>
        )}
      </div>
    </>
  )
}

export default App

