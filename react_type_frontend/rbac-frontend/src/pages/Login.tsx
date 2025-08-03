import { useForm } from 'react-hook-form';
import { useDispatch } from 'react-redux';
import type { AppDispatch } from '../app/store';
import { loginSuccess } from '../features/auth/authSlice';
import './Auth.css';

interface FormData {
  username: string;
  password: string;
}

const Login = () => {
  const { register, handleSubmit } = useForm<FormData>();
  const dispatch = useDispatch<AppDispatch>();

  const onSubmit = async (data: FormData) => {
    const res = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });

    if (res.ok) {
      const result = await res.json();
      dispatch(loginSuccess(result.token));
      alert('Login successful');
    } else {
      alert('Invalid credentials');
    }
  };

  return (
    <form className="auth-form" onSubmit={handleSubmit(onSubmit)}>
      <h2>Login</h2>
      <input {...register('username')} placeholder="Username" required />
      <input {...register('password')} type="password" placeholder="Password" required />
      <button type="submit">Login</button>
    </form>
  );
};

export default Login;
