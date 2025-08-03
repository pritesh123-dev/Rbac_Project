import { useSelector } from 'react-redux';
import type { RootState } from '../app/store';

const Dashboard = () => {
  const user = useSelector((state: RootState) => state.auth.user);

  return (
    <div>
      <h2>Dashboard</h2>
      <p>Welcome, <strong>{user?.username}</strong>!</p>
      <p>Your role: <strong>{user?.role}</strong></p>
    </div>
  );
};

export default Dashboard;
