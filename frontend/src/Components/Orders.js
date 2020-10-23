import React from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Title from './Title';
import CheckBoxRoundedIcon from '@material-ui/icons/CheckBoxRounded';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import DeleteIcon from '@material-ui/icons/Delete';
import ListItem from '@material-ui/core/ListItem';
import axios from 'axios';
import {
  Link
} from "react-router-dom";

export default class Orders extends React.Component {
  constructor(props) {
    super(props);
    this.state = { records: [] };

    this.Availablity = this.Availablity.bind(this);
  }
  //ws = new WebSocket('ws://localhost/status')
  componentDidMount() {
    //const protocolPrefix = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
   // let { host } = window.location; // nb: window location contains the port, so host will be localhost:3000 in dev
    //this.ws = new WebSocket(`${protocolPrefix}//${host}/ws/status`);
    //this.ws = new WebSocket('ws://localhost/ws/status');
    this.ws = new WebSocket('ws://localhost:80/ws/status/');

    this.ws.onopen = () => {
      // on connecting, do nothing but log it to the console
      console.log('connected')
      this.ws.send('')
    }

    this.ws.onmessage = evt => {
      // listen to data sent from the websocket server
      const message = JSON.parse(evt.data)
      this.setState({ records: message });
      console.log(evt.data)
    }

    this.ws.onclose = () => {
      console.log('disconnected')
      // automatically try to reconnect on connection loss

    }

  }
  componentWillUnmount() {
    this.ws.close();
    console.log('unmounted')
  }
  Availablity(row) {
    if (row.desired === row.lastCode) {
      return <CheckBoxRoundedIcon style={{ color: "green" }} />
    } else {
      return <CancelRoundedIcon style={{ color: "red" }} />
    }
  }

  render() {
    return (
      <React.Fragment>
        <Title>Status Table</Title>
        <Table size="small">
          <TableHead>
            <TableRow>
              <TableCell><b>Availablity</b></TableCell>
              <TableCell><b>Address</b></TableCell>
              <TableCell><b>Method</b></TableCell>
              <TableCell><b>Desired Status Code</b></TableCell>
              <TableCell><b>Proxy</b></TableCell>
              <TableCell><b>Checking Interval</b></TableCell>
              <TableCell><b>Status</b></TableCell>
              <TableCell align="right"><b>Action</b></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {this.state.records.map((row) => (
              <TableRow key={row.id}>
                <TableCell>{this.Availablity(row)}</TableCell>
                <TableCell>{row.host}</TableCell>
                <TableCell>{row.method}</TableCell>
                <TableCell>{row.desired}</TableCell>
                <TableCell>{row.proxy}</TableCell>
                <TableCell>{row.interval}</TableCell>
                <TableCell>{row.lastCode}</TableCell>
                <TableCell align="right"><Delete row={row} /></TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </React.Fragment>
    );
  }
}
class Delete extends React.Component {
  constructor(props) {
    super(props);
    this.state = { row: this.props.row };

    this.DeleteRow = this.DeleteRow.bind(this);
  }
  DeleteRow() {

    const headers = {
      'Content-Type': 'application/json'
    };

    axios.post(
      '/backend/delete/',
      {
        host: this.props.row.host,
        desired: this.props.row.desired,
        interval:  this.props.row.interval,
        method: this.props.row.method,
        proxy: this.props.row.proxy,
        lastCode: this.props.row.desired,
      },
      { headers }
    )
      .then(response => { console.log(response.data) })
      .catch(error => { console.log("Error ========>", error); }
      )
  }
  render() {
    return (
      <ListItem button component={Link} to={'/redirect'} onClick={this.DeleteRow}>
        <DeleteIcon />
      </ListItem>
    );
  }
}