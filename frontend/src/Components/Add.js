import React from 'react';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControl from '@material-ui/core/FormControl';
import FormHelperText from '@material-ui/core/FormHelperText';
import NativeSelect from '@material-ui/core/NativeSelect';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';


const useStyles = makeStyles((theme) => ({
    paper: {
        marginTop: theme.spacing(8),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    avatar: {
        margin: theme.spacing(1),
        backgroundColor: theme.palette.secondary.main,
    },
    form: {
        width: '100%', // Fix IE 11 issue.
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
    formControl: {
        margin: theme.spacing(1),
        minWidth: 120,
    },
    selectEmpty: {
        marginTop: theme.spacing(2),
    },
}));

export default function Add() {
    const classes = useStyles();

    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <div className={classes.paper}>
                <Typography component="h1" variant="h5">
                    Add new target
        </Typography>
                <form className={classes.form} noValidate>
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        id="target"
                        label="Target Address"
                        type="text"
                        name="target"
                        autoFocus
                    />
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        name="proxy"
                        label="Proxy"
                        type="proxy"
                        id="proxy"
                        autoFocus
                    />
                    <FormControl className={classes.formControl}>
                        <NativeSelect
                            className={classes.selectEmpty}
                            value="10"
                            name="age"
                        >
                            <option value={10}>GET</option>
                            <option value={20}>POST</option>
                        </NativeSelect>
                        <FormHelperText>Method</FormHelperText>
                    </FormControl>
                    <FormControl className={classes.formControl}>
                        <NativeSelect
                            className={classes.selectEmpty}
                            value="200"
                            name="Desired Status Code"
                        >
                            <option value={100}>100</option>
                            <option value={200}>200</option>
                            <option value={300}>300</option>
                            <option value={400}>400</option>
                            <option value={500}>500</option>
                        </NativeSelect>
                        <FormHelperText>Desired Status Code</FormHelperText>
                    </FormControl>
                    <FormControl className={classes.formControl}>
                        <NativeSelect
                            className={classes.selectEmpty}
                            value="200"
                            name="Checking time interval"
                        >
                            <option value={10}>10s</option>
                            <option value={30}>30s</option>
                            <option value={60}>1m</option>
                            <option value={90}>1.5m</option>
                            <option value={120}>2m</option>
                            <option value={300}>5m</option>
                            <option value={600}>10m</option>
                            <option value={1800}>30m</option>
                            <option value={3600}>m</option>
                        </NativeSelect>
                        <FormHelperText>Checking time interval</FormHelperText>
                    </FormControl>
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        color="primary"
                        className={classes.submit}
                    >
                        Add
          </Button>

                </form>
            </div>
        </Container>
    );
}