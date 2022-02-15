import React from 'react'
import { BrowserRouter as Router, Navigate  ,Route, Routes } from "react-router-dom";
import Calculadora from "./Componentes/Calculadora.js"
import TablaRegistros from "./Componentes/TablaRegistros"

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/TablaRegistros" element={<TablaRegistros />} />
        <Route path="/Calculadora" element={<Calculadora />} />
        <Route path="/" element={<Navigate replace to="/Calculadora" />} />
      </Routes>
    </Router>
  );
}

export default App;
