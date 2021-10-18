import React from "react";
import Nbody from "./components/AppBody";
import test from "./components/test";
import Nhead from "./components/AppHeader";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";

function App() {
  return (
    <Router>
    <div>
        <Nhead/>
        
        <Switch>

        
        <Route exact path="/" component={Nbody} />
        <Route exact path="/test" component={test} />

       </Switch>

    </div>
    </Router>
  );
}
///<Route exact path="/บันทึก" component={test} />
export default App;

