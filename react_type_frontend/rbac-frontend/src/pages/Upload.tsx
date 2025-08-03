import { useRef } from 'react';
import type { FormEvent } from 'react';
import { useSelector } from 'react-redux';
import type { RootState } from '../app/store';

const Upload = () => {
  const fileRef = useRef<HTMLInputElement>(null);
  const token = useSelector((state: RootState) => state.auth.token);

  const handleUpload = async (e: FormEvent) => {
    e.preventDefault();
    if (!fileRef.current?.files?.[0]) return;

    const formData = new FormData();
    formData.append('file', fileRef.current.files[0]);

    const res = await fetch('http://localhost:8080/api/upload', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: formData,
    });

    if (res.ok) {
      alert('Upload successful');
    } else {
      alert('Upload failed');
    }
  };

  return (
    <form onSubmit={handleUpload}>
      <h2>Upload Document</h2>
      <input type="file" ref={fileRef} required />
      <button type="submit">Upload</button>
    </form>
  );
};

export default Upload;
