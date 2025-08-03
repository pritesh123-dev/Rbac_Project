// src/routes/PrivateRoute.tsx
import { useSelector } from 'react-redux';
import { Navigate } from 'react-router-dom';
import type { RootState } from '../app/store';
import type { JSX } from 'react';


type Props = {
  children: JSX.Element;
  roles?: string[];
};

const PrivateRoute = ({ children, roles }: Props) => {
  const { token, user } = useSelector((state: RootState) => state.auth);

  if (!token || !user) return <Navigate to="/login" replace />;
  if (roles && !roles.includes(user.role)) return <Navigate to="/unauthorized" replace />;

  return children;
};

export default PrivateRoute;
