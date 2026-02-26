function MyRentals({ rentals, onBack }) {
  return (
    <div className="my-rentals">
      {rentals.length === 0 ? (
        <p className="no-rentals">No active rentals.</p>
      ) : (
        <div className="rentals-grid">
          {rentals.map((rental) => (
            <div key={rental.public_id} className="rental-card">
              <img
                className="rental-cover"
                src={`/tapes/${rental.tape_title}.png`}
              />
              <h4 className="rental-title">{rental.tape_title}</h4>
              <p className="rental-date">
                Rented: {new Date(rental.rented_at).toLocaleDateString()}
              </p>
            </div>
          ))}
        </div>
      )}
      <button className="panel-button" onClick={onBack}>
        Back
      </button>
    </div>
  );
}

export default MyRentals;
