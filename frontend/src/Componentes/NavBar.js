import React,{useEffect} from 'react'
import Box from '@mui/material/Box';
import BottomNavigation from '@mui/material/BottomNavigation';
import BottomNavigationAction from '@mui/material/BottomNavigationAction';
import RestoreIcon from '@mui/icons-material/Restore';
import FavoriteIcon from '@mui/icons-material/Favorite';
import { useNavigate  } from 'react-router-dom';

export default function NavBar(props) {
  const [value, setValue] = React.useState(0);
  const navigate = useNavigate();

  useEffect(() => {
    async function obtener() {
        if (props.value === "Calculadora"){
          setValue(0)
        }else{
          setValue(1)
        }
    }
    obtener()
    });

  const CambiarPestaña = (newValue) =>{
    if (newValue === 0) {
        navigate('/Calculadora');
        setValue(0);
    }else if (newValue === 1){
        navigate('/TablaRegistros');
        setValue(1);
    }
    return newValue
  }

  return (
    <Box sx={{ mx: 'auto', width: 400 }}>
      <BottomNavigation
        showLabels
        value={value}
        onChange={(event, newValue) => {
          setValue(CambiarPestaña(newValue));   
        }}
        sx={{ bgcolor: '#848482' }}
      >
        <BottomNavigationAction label="Calculadora" icon={<FavoriteIcon />} />
        <BottomNavigationAction label="Registro De Operaciones" icon={<RestoreIcon />} />
      </BottomNavigation>
    </Box>
  );
}