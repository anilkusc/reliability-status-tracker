import React from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import DashboardIcon from '@material-ui/icons/Dashboard';
import AddIcon from '@material-ui/icons/Add';
import {
  Link
} from "react-router-dom";


export const mainListItems = (
  <div>
    <ListItem button component={Link} to={'/'}>
      <ListItemIcon>
        <DashboardIcon />
      </ListItemIcon>
      <ListItemText primary="Status" />
    </ListItem>
    <ListItem button component={Link} to={'/add'}>  
      <ListItemIcon>
        <AddIcon />
      </ListItemIcon>

      <ListItemText primary="Add"  />
    </ListItem>
  </div>
);