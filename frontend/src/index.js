import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import Root from "./Root";
import * as serviceWorkerRegistration from "./serviceWorkerRegistration";
import reportWebVitals from "./reportWebVitals";

export const RenderApp = () =>
  ReactDOM.render(
    <React.StrictMode>
      <Root />
    </React.StrictMode>,
    document.getElementById("root")
  );
RenderApp();
serviceWorkerRegistration.unregister();
reportWebVitals();
