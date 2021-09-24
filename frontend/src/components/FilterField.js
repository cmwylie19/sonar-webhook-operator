import * as React from 'react';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import { makeStyles } from '@mui/styles';
import { green } from "@material-ui/core/colors";

const useStyles = makeStyles({
    root: {
      "& .css-cio0x1-MuiInputBase-root-MuiFilledInput-root:after": {
          borderColor: "green"
      },
      "& .css-10botns-MuiInputBase-input-MuiFilledInput-input":{
          color: "white"
      },
      "& .MuiFormLabel-colorPrimary":{
          color: "white"
      },

      "& .Mui-focused": {
        fontColor: "tomato", 
        fontWeight: "bold"
      },
      "& label.Mui-focused": {
        color: "lime"
      },
    }
  });

export default function FilterField() {
    const classes = useStyles()
  return (
    <Box
      component="form"
      sx={{
        '& > :not(style)': { m: 1, width: '100%' },
      }}
      noValidate
      autoComplete="off"
    >
      <TextField className={classes.root} id="outlined-basic" label="Filter results" variant="filled" />
    </Box>
  );
}