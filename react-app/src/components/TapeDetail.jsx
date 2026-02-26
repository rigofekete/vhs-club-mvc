function TapeDetail({ tape, error, onBack, onRent }) {
  return (
    <div className="tape-detail">
      <img className="tape-detail-cover"
        src={`/tapes/${tape.title}.png`}
      />
      <h3>{tape.title}</h3>
      <p>{tape.director}</p>
      {error && <p className="error-message">{error}</p>}
      <button className="panel-button" onClick={onRent}>
        Rent
      </button>
      <button className="panel-button" onClick={onBack}>
        Back
      </button>
    </div>
  );
}

export default TapeDetail;
