// Función para procesar y validar el correo electrónico

function validateNotEmpty(email) {
  if (!email) {
    throw new Error("El correo electrónico no puede estar vacío.");
  }
}

function validateMaxLength(email) {
  if (email.length > 150) {
    throw new Error("El correo electrónico no debe superar los 150 caracteres.");
  }
}

function validateFormat(email) {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!emailRegex.test(email)) {
    throw new Error("El formato del correo electrónico no es válido.");
  }
}

function sanitizeEmail(email) {
  return email.replace(/[&<>"']/g, (char) => {
    const entities = {
      "&": "&amp;",
      "<": "&lt;",
      ">": "&gt;",
      '"': "&quot;",
      "'": "&#039;"
    };
    return entities[char];
  });
}

export function processEmail(email) {
  validateNotEmpty(email);
  email = email.trim();
  validateMaxLength(email);
  validateFormat(email);
  return sanitizeEmail(email);
}