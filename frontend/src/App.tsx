import React, { useEffect } from "react";
import Nbody from "./components/AppBody";
import test from "./components/test";
import Nhead from "./components/AppHeader";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import SignIn from "./components/SignIn";

function App() {
  const [token, setToken] = React.useState<String>("");

  useEffect(() => {
    const getToken = localStorage.getItem("uid");
    if (getToken) {
      setToken(getToken);
    }
  }, []);

  if (!token) {
    return <SignIn />
  }
  
  return (
    <Router>
      <div>
        <Nhead />

        <Switch>
        <Route exact path="/" component={test} />
          <Route exact path="/Nbody" component={Nbody} />
          
        </Switch>

      </div>
    </Router>
  );
}
///<Route exact path="/บันทึก" component={test} />
export default App;

///
/*
<Route exact path="/" component={Nbody} />
        <Route exact path="/test" component={test} />
  */