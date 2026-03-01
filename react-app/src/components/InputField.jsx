
function InputField(props) {
  const className = props.className;
  const type = props.type;
  const label = props.label;
  //TODO: Any way to get diagnostic warnings in jsx? for example if I type Onchange here, there is no lsp warning for undefined variable.
  const onChange = props.onChange;

  return (
    <div className={`input-wrapper ${className}`}>
      <label>
        {/* {label} */}
      </label>
      {/* <input type={type} onChange={onChange} /> */}
      <input type={type} onChange={onChange} placeholder={label} />
    </div>
  );
}



export default InputField;

