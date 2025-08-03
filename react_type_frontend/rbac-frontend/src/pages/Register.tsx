import { useForm } from 'react-hook-form';
import './Auth.css';

interface FormData {
  username: string;
  password: string;
}

const Register = () => {
  const { register, handleSubmit } = useForm<FormData>();

  const onSubmit = async (data: FormData) => {
    const res = await fetch('http://localhost:8080/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });

    if (res.ok) {
      alert('Registration successful');
    } else {
      alert('Failed to register');
    }
  };

  return (
    <form className="auth-form" onSubmit={handleSubmit(onSubmit)}>
      <h2>Register</h2>
      <input {...register('username')} placeholder="Username" required />
      <input {...register('password')} type="password" placeholder="Password" required />
      <button type="submit">Register</button>
    </form>
  );
};

export default Register;
