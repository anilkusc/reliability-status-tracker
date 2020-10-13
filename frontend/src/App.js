import React from 'react';
import Dashboard from './Components/Dashboard'
import SignIn from './Components/SignIn'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
  Link
} from "react-router-dom";

function App() {

  return (
    <Routes />
  );
}

class Routes extends React.Component {
  constructor(props) {
    super(props);
    this.state = { isLoggedIn: false };
    this.handleSetLoggedIn = this.handleSetLoggedIn.bind(this);
  }

  handleSetLoggedIn() {
    this.setState({ isLoggedIn: true });
  }

  render() {
    if (this.state.isLoggedIn) {
      return (
        <Router>
          <Switch>
            <Route exact path="/" >
              <Dashboard content="orders" />
            </Route>
            <Route exact path="/redirect" >
            <Redirect  from="/redirect"  to="/"  exact  />
            </Route>
            <Route exact path="/add" >
              <Dashboard content="add" />
            </Route>
            <Redirect  from="/signin"  to="/"  exact  />
          </Switch>
        </Router>
      );
    } else {
      return (
        <SignIn handleSetLoggedIn={this.handleSetLoggedIn}/>
      );
    }


  }
}
export default App;
