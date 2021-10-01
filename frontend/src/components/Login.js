import React, { useEffect, useState } from "react";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import CssBaseline from "@mui/material/CssBaseline";
import TextField from "@mui/material/TextField";
import FormControlLabel from "@mui/material/FormControlLabel";
import Checkbox from "@mui/material/Checkbox";
import Box from "@mui/material/Box";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { makeStyles } from "@mui/styles";
import { green, grey } from "@material-ui/core/colors";
import Link from "@mui/material/Link";
import Grid from "@mui/material/Grid";
import API from "../utils";
const theme = createTheme({
  palette: {
    mode: "dark",
    primary: {
      main: grey[900],
    },
    secondary: { main: green[400] },
  },
});

const useStyles = makeStyles({
  root: {
    "& .css-cio0x1-MuiInputBase-root-MuiFilledInput-root:after": {
      borderColor: "green",
    },
    "& .css-10botns-MuiInputBase-input-MuiFilledInput-input": {
      color: "white",
    },
    "& .MuiFormLabel-colorPrimary": {
      color: "white",
    },

    "& .Mui-focused": {
      fontColor: "tomato",
      fontWeight: "bold",
    },
    "& label.Mui-focused": {
      color: "lime",
    },
  },
});

export default function Login() {
  useEffect(() => {
    // attempt to login user
    // if successful, push to app
  }, []);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const classes = useStyles();
  const handleSubmit = async (event) => {
    event.preventDefault();
    alert("email: " + email + " password: " + password);
    const api = new API();
    console.log({
      email: email,
      password: password,
    });
    let response = await api.authenticateUser(email, password);

    alert(JSON.stringify(response, undefined, 2));
    // eslint-disable-next-line no-console
  };

  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          <Typography component="p" variant="p" id="description">
            Whats the deal
          </Typography>
          <Box
            component="form"
            onSubmit={handleSubmit}
            noValidate
            sx={{ mt: 1 }}
          >
            <TextField
              inputProps={{
                // autocomplete: "email",
                form: {
                  autocomplete: "off",
                },
              }}
              margin="normal"
              className={classes.root}
              id="email"
              label="Email Address"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              name="email"
              fullWidth
              variant="filled"
              autoFocus
            />
            <TextField
              inputProps={{
                autocomplete: "password",
                form: {
                  autocomplete: "off",
                },
              }}
              margin="normal"
              className={classes.root}
              name="password"
              label="Password"
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              fullWidth
              variant="filled"
              autoFocus
            />

            <FormControlLabel
              control={
                <Checkbox
                  value="remember"
                  color="primary"
                  name="remember"
                  id="remember"
                  color="success"
                />
              }
              label="Remember me"
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2, bgcolor: "primary.light" }}
            >
              Sign In
            </Button>
          </Box>
          <Grid container justifyContent="center">
            <Grid item>
              <Link
                href="/signup"
                variant="body2"
                sx={{ mt: 3, mb: 2, color: "secondary.dark" }}
              >
                Don't have an account? Signup
              </Link>
            </Grid>
          </Grid>
        </Box>
      </Container>
    </ThemeProvider>
  );
}
