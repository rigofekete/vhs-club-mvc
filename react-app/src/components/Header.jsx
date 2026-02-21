// Using props object, automatically created by React when passed to the parent in App.jsx
// this props can contain multiple accessible attributes (e.g. age, id)
// function Header(props) {
//   const name = props.name;
//   return <h1>Welcome, {name}!</h1>;
// }

// function Header() {
//   // return <h1 >VHS CLUB</h1>;
//   return <div className="header" ><h1>VHS CLUB</h1></div>;
// }

function Header(props) {
  const logged = props.logged;
  return (
    <pre className={logged ? "header header-blink" : "header"}>{String.raw`██╗   ██╗  ██╗  ██╗  ███████╗      ██████╗  ██╗      ██╗   ██╗  ██████╗ 
██║   ██║  ██║  ██║  ██╔════╝     ██╔════╝  ██║      ██║   ██║  ██╔══██╗
██║   ██║  ███████║  ███████╗     ██║       ██║      ██║   ██║  ██████╔╝
╚██╗ ██╔╝  ██╔══██║  ╚════██║     ██║       ██║      ██║   ██║  ██╔══██╗
 ╚████╔╝   ██║  ██║  ███████║     ╚██████╗  ███████╗ ╚██████╔╝  ██████╔╝
 ╚═══╝    ╚═╝  ╚═╝  ╚══════╝      ╚═════╝  ╚══════╝  ╚═════╝   ╚═════╝`}</pre>
  );
}

// function Header(props) {
//   const logged = props.logged;
//   return (
//     <h1 className={logged ? "header header-blink" : "header"}>
//       VHS CLUB
//     </h1>
//   );
// }

//Alternative syntax, destructuring the name attribute from the props object:
// function Header({ name }) {
//   return <h1>Welcome, {name}!</h1>;
// }

export default Header;
