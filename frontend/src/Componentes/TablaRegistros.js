import React,{useEffect} from 'react'
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TablePagination from '@mui/material/TablePagination';
import TableRow from '@mui/material/TableRow';
import NavBar from '../Componentes/NavBar.js';
import CssBaseline from '@mui/material/CssBaseline';
import '../css/Calculadora.css';
import axios from "axios";

const columns = [
  { id: 'Primero', label: 'Numero 1', minWidth: 100 },
  { id: 'Segundo', label: 'Numero 2', minWidth: 100 },
  { id: 'Operacion', label: 'Operacion', minWidth: 100 },
  { id: 'Resultado', label: 'Resultado', minWidth: 200 },
  { id: 'Fecha', label: 'Fecha Y Hora', minWidth: 200 },
];

export default function TablaRegistros() {
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const [rows, setRows] = React.useState([])

  useEffect(() => {
    async function obtener() {
        const data = await axios.get("http://localhost:3000/RegresarOperaciones");
        setRows(data.data.Operaciones)
    }
    obtener()
    },[]);

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setRowsPerPage(+event.target.value);
    setPage(0);
  };

  return (
    <div className='Fondo'>
    <Paper sx={{ width: '100%', overflow: 'hidden', backgroundColor: '#474b4e' }}>
    <CssBaseline />
    <NavBar value="Tabla"/>
      <TableContainer sx={{ maxHeight: 440, backgroundColor: '#8f8f8f' }}>
        <Table aria-label="sticky table">
          <TableHead sx={{ bgcolor: '#686868' }}>
            <TableRow>
              {columns.map((column) => (
                <TableCell
                  key={column.id}
                  align={column.align}
                  style={{ minWidth: column.minWidth }}
                >
                  {column.label}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {rows
              .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
              .map((row) => {
                return (
                  <TableRow hover role="checkbox" tabIndex={-1} key={row.code}>
                    {columns.map((column) => {
                      const value = row[column.id];
                      return (
                        <TableCell key={column.id} align={column.align}>
                          {column.format && typeof value === 'number'
                            ? column.format(value)
                            : value}
                        </TableCell>
                      );
                    })}
                  </TableRow>
                );
              })}
          </TableBody>
        </Table>
      </TableContainer>
      <TablePagination
        rowsPerPageOptions={[10, 25, 100]}
        component="div"
        count={rows.length}
        rowsPerPage={rowsPerPage}
        page={page}
        onPageChange={handleChangePage}
        onRowsPerPageChange={handleChangeRowsPerPage}
      />
    </Paper>
    </div>
  );
}