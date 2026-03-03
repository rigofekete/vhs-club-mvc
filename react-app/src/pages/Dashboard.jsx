import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import './Dashboard.css'
import TapeDetail from '../components/TapeDetail'
import MyRentals from '../components/MyRentals'
import ParabolicBackground from '../components/ParabolicBackground'



function Dashboard() {
  const navigate = useNavigate();
  const username = localStorage.getItem('username')
  const token = localStorage.getItem('token')
  const [view, setView] = useState('catalog');
  const [videoFadingOut, setVideoFadingOut] = useState(false);
  const [selectedTape, setSelectedTape] = useState(null);
  const [animKey, setAnimKey] = useState(0);
  const [tapes, setTapes] = useState([]);
  const [rentals, setRentals] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!token) {
      navigate('/login');
      return;
    }
  }, [token, navigate])

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080/tapes');
        if (!response.ok) {
          const errData = await response.json();
          let errorMsg = errData.error || 'Failed to fetch tapes';
          if (errData.fields) {
            errorMsg += ': ' + Object.entries(errData.fields)
              .map(([field, msg]) => `${field} - ${msg}`)
              .join(', ');
          }
          throw errorMsg;
        }

        const data = await response.json();
        setTapes(data);
      } catch (err) {
        setError(err)
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  useEffect(() => {
    if (view === 'thankYou') {
      const timer = setTimeout(() => {
        fetchRentals();
        setView('myRentals');
        setSelectedTape(null);
        setAnimKey(prev => prev + 1);
      }, 2000);
      return () => clearTimeout(timer);
    }
  }, [view]);

  const fetchRentals = async () => {
    try {
      const response = await fetch('http://localhost:8080/rentals');
      if (!response.ok) return;
      const data = await response.json();
      setRentals(data.filter(r => r.username === username));
    } catch (err) {
      setError('Unable to fetch rentals.');
    }
  };

  const handleLogout = () => {
    localStorage.removeItem('token')
    localStorage.removeItem('username')
    navigate('/login');
  };

  const handleCatalog = () => {
    setView('catalog');
    setSelectedTape(null);
    setError('');
    setAnimKey(prev => prev + 1);
  };

  const handleMyRentals = () => {
    fetchRentals();
    setView('myRentals');
    setSelectedTape(null);
    setError('');
    setAnimKey(prev => prev + 1);
  };

  const handleTapeClick = (tape) => {
    setSelectedTape(tape);
    setView('detail');
    setAnimKey(prev => prev + 1);
  };

  const handleRent = async (tape) => {
    try {
      const response = await fetch(`http://localhost:8080/rentals/${tape.public_id}`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        const errData = await response.json();
        setError(errData.error || 'Failed to rent tape');
        return;
      }

      setView('renting');
    } catch (err) {
      setError('Unable to connect to server.');
    }
  };


  return (
    <div className="Dashboard" >
      <ParabolicBackground animateDuration={1000} trigger={animKey} />
      <div className="dashboard-content">
        <aside className="user-panel">
          <img className="user-photo" src={`/profile/${username}.png`} />
          <h2 className="user-name">{username}</h2>
          <button className="panel-button" onClick={handleMyRentals}>My Rentals</button>
          <button className="panel-button" onClick={handleCatalog}>
            Catalog
          </button>
          <button className="panel-button logout-button" onClick={handleLogout}>
            Logout
          </button>
        </aside>

        <div className="tape-catalog" key={animKey}>
          <div className="catalog-fade">
            <h2>{view === 'myRentals' ? 'My Rentals' : 'VHS Catalog'}</h2>

            {/* work on this, it looks choppy */}
            {view === 'renting' && (
              <div className="rent-animation-wrapper">
                <video
                  className={`rent-animation ${videoFadingOut ? 'fade-out' : ''}`}
                  src="/videos/rented.mp4"
                  autoPlay
                  onEnded={() => {
                    setError('');
                    setVideoFadingOut(true);
                    setTimeout(() => {
                      setVideoFadingOut(false);
                      setView('thankYou');
                    }, 1000);
                  }}
                />
              </div>
            )}

            {view === 'thankYou' && (
              <p className="thank-you-text">Thank you for renting!</p>
            )}

            {view === 'myRentals' && (
              <MyRentals
                rentals={rentals}
                onBack={handleCatalog}
              />
            )}

            {view === 'detail' && selectedTape && (
              <TapeDetail
                tape={selectedTape}
                error={error}
                onBack={handleCatalog}
                onRent={() => handleRent(selectedTape)}
              />
            )}

            {view === 'catalog' && (
              <>
                {loading && <p className="loading-text">Loading tapes...</p>}
                {error && <p className="error-message">{error}</p>}

                <div className="tape-grid">
                  {tapes.map((tape) => (
                    <div key={tape.public_id} className="tape-card" onClick={() => handleTapeClick(tape)}>
                      <img className="tape-cover"
                        src={`/tapes/${tape.title}.png`}
                      />
                      <h4 className="tape-title">{tape.title}</h4>
                      <p className="tape-director">{tape.director}</p>
                    </div>
                  ))}
                </div>
              </>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

export default Dashboard
