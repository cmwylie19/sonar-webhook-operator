import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import Root from "./Root";
import Login from "./components/Login";
import Signup from "./components/Signup";
import * as serviceWorkerRegistration from "./serviceWorkerRegistration";
import reportWebVitals from "./reportWebVitals";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

export const RenderApp = () =>
  ReactDOM.render(
    <React.StrictMode>
      <Router>
        <Switch>
          <Route path="/login" render={() => <Login />} />
          <Route path="/signup" render={() => <Signup />} />
          <Route path="/" render={() => <Root />} />
        </Switch>
      </Router>
    </React.StrictMode>,
    document.getElementById("root")
  );
RenderApp();
serviceWorkerRegistration.unregister();
reportWebVitals();
