import { createSlice } from '@reduxjs/toolkit';
import type { PayloadAction } from '@reduxjs/toolkit';
import { jwtDecode } from 'jwt-decode';

interface JWTClaim {
  username: string;
  role: string;
  exp: number;
}

interface AuthState {
  token: string | null;
  user: JWTClaim | null;
}

const token = localStorage.getItem('token');
const decoded = token ? jwtDecode<JWTClaim>(token) : null;

const initialState: AuthState = {
  token,
  user: decoded,
};

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    loginSuccess: (state, action: PayloadAction<string>) => {
      const decoded = jwtDecode<JWTClaim>(action.payload);
      state.token = action.payload;
      state.user = decoded;
      localStorage.setItem('token', action.payload);
    },
    logout: (state) => {
      state.token = null;
      state.user = null;
      localStorage.removeItem('token');
    },
  },
});

export const { loginSuccess, logout } = authSlice.actions;
export default authSlice.reducer;
