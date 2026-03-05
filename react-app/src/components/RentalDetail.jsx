function RentalDetail({ rental, error, onBack, onReturnRent }) {
  return (
    <div className="tape-detail">
      <img className="tape-detail-cover"
        src={`/tapes/${rental.tape_title}.png`}
      />
      <h3>{rental.tape_title}</h3>
      <p className="rental-date">
        Rented at: {new Date(rental.rented_at).toDateString()}
      </p>
      <button className="panel-button" onClick={onReturnRent}>
        Return Tape
      </button>
      <button className="panel-button" onClick={onBack}>
        Back
      </button>
      {error && <h1 className="error-message">{error}</h1>}
    </div>
  );
}

export default RentalDetail;
