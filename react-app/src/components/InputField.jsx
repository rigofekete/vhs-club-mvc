
function InputField(props) {
  const className = props.className;
  const type = props.type;
  const label = props.label;
  const onChange = props.onChange;

  return (
    <div className={`input-wrapper ${className}`}>
      <input type={type} onChange={onChange} placeholder={label} />
    </div>
  );
}



export default InputField;

