import React from 'react'
import {Switch, BrowserRouter as Router, Route} from "react-router-dom"
import Login from "./components/login"
import LoginCallback from "./components/loginCallback"

import doc from "./oauth-config.json"
const App = () => {

  const google = doc.filter(d => d.provider==="google")[0]
  const facebook = doc.filter(d => d.provider==="facebook")[0]
  const github = doc.filter(d => d.provider==="github")[0]


  return (
  <Router>
    <Switch>
    <Route path="/login" component={Login} />
    <Route path="/" exact component={Login} />
    <Route path="/google/callback" component={(props) => <LoginCallback  {...props} uri={google.backend_url+"/callback"}/>}  />
    <Route path="/facebook/callback" component={(props) => <LoginCallback  {...props} uri={facebook.backend_url+"/callback"}/>}  />
    <Route path="/github/callback" component={(props) => <LoginCallback  {...props} uri={github.backend_url+"/callback"}/>}  />
    </Switch>
  </Router>
  )
}

export default App
