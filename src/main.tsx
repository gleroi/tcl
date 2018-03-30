import * as React from "react";
import * as ReactDOM from "react-dom";

import App from "./app";
import "./tcl.scss";

ReactDOM.render(<App/>, document.querySelector("#root"));

navigator.serviceWorker.register('/sw.js')
.then(registration => {
    console.log("SW registration: OK", registration)
})
.catch(err => {
    console.error("SW registration, KO: " + err);
});

