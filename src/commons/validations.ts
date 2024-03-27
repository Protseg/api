export function isEmailValid(email: string): boolean {
  const emailRegex: RegExp = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i;
  return emailRegex.test(email);
}

export function isPasswordValid(password: string): boolean {
  if (!password) {
    return false;
  }

  if (password.length < 10) {
    return false;
  }

  // At leasr 3 different characters
  if (new Set(password).size < 3) {
    return false;
  }

  return true;
}
