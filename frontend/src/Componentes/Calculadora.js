import * as React from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import Container from '@mui/material/Container';
import TextField from '@mui/material/TextField';
import MenuItem from '@mui/material/MenuItem';
import axios from "axios"
import '../css/Calculadora.css'

const operaciones = [
  {
    value: '+',
    label: '+',
  },
  {
    value: '-',
    label: '-',
  },
  {
    value: '*',
    label: '*',
  },
  {
    value: '/',
    label: '/',
  },
];

function Calculadora() {
    const [primero, setPrimero] = React.useState('0');
    const [segundo, setSegundo] = React.useState('0');
    const [operacion, setOperacion] = React.useState('+');
  
    const SeleccionOperacion = (event) => {
      setOperacion(event.target.value);
    };
  
    const CambiarPrimero = (event) => {
      setPrimero(event.target.value);
    };
  
    const CambiarSegundo = (event) => {
      setSegundo(event.target.value);
    };
  
    const Operar = async()=>{
      const Valores={
        "Primero": parseFloat(primero),
        "Segundo": parseFloat(segundo),
        "Operacion": operacion
      }
      console.log(Valores)
      await axios.post("http://localhost:3000/ObtenerOperacion",Valores)
    }

    return ( 
        <div className='Fondo'>
          <React.Fragment>
          <CssBaseline />
          <Container maxWidth="sm">
            <br/><br/><br/><br/><br/><br/><br/><br/><br/>
            <Box sx={{ bgcolor: '#828282', width:300, heigth:600, mx: 'auto', p:2}} > 
              <TextField id="primero" color="success" helperText="Ingrese Primer Numero" value={primero} onChange={CambiarPrimero} focused/><br/><br/>
    
              <TextField id="segundo" color="success" helperText="Ingrese Segundo Numero" value={segundo} onChange={CambiarSegundo} focused/><br/><br/>
    
              <TextField id="operacion" select label="Select" value={operacion} onChange={SeleccionOperacion} color="warning" helperText="Seleccion Operacion">
                {operaciones.map((option) => ( <MenuItem key={option.value} value={option.value}> {option.label} </MenuItem> ))}
              </TextField><br/><br/>
    
              <Button id="Enviar" variant="contained" color="primary" onClick={Operar}>Operar</Button>
            </Box>
          </Container>
        </React.Fragment>
          
        </div>
    );
}

export default Calculadora;