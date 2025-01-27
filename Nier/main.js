// Importar la función desde commons.js
import { processEmail } from "./commons.js";

// Función principal
function main() {
  const emailInput = prompt("Por favor, ingresa tu correo electrónico:");
  
  try {
    const sanitizedEmail = processEmail(emailInput);
    console.log("Correo válido:", sanitizedEmail);
  } catch (error) {
    console.error("Error:", error.message);
  }
}

// Ejecutar el programa principal
main();