// src/components/Navbar.tsx
import { Link } from 'react-router-dom';
import { useSelector } from 'react-redux';
import type { RootState } from '../app/store';

const Navbar = () => {
  const { user } = useSelector((state: RootState) => state.auth);

  return (
    <nav style={{ padding: '10px', background: '#eee' }}>
      <Link to="/dashboard">Dashboard</Link>

      {user?.role === 'admin' && (
        <Link to="/admin" style={{ marginLeft: '10px' }}>
          Admin Panel
        </Link>
      )}

      {user?.role === 'editor' && (
        <Link to="/upload" style={{ marginLeft: '10px' }}>
          Upload
        </Link>
      )}

      {!user && (
        <>
          <Link to="/login" style={{ marginLeft: '10px' }}>Login</Link>
          <Link to="/register" style={{ marginLeft: '10px' }}>Register</Link>
        </>
      )}
    </nav>
  );
};

export default Navbar;
