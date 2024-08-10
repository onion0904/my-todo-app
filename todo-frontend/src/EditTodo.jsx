import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useNavigate, useParams } from 'react-router-dom';

function EditTodo() {
    const { id } = useParams();
    const [todo, setTodo] = useState({ Name: '', Text: '', TimeLimit: '' });
    const navigate = useNavigate();

    useEffect(() => {
        axios.get(`http://localhost:8080/Todo/get/${id}`)
            .then(response => {
                setTodo({
                    Name: response.data.Name,
                    Text: response.data.Text,
                    TimeLimit: response.data.TimeLimit
                });
            })
            .catch(error => {
                console.error('Error fetching todo:', error);
            });
    }, [id]);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setTodo(prev => ({
            ...prev,
            [name]: value
        }));
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        axios.put(`http://localhost:8080/Todo/update/${id}`, todo)
            .then(() => {
                navigate('/');
            })
            .catch(error => {
                console.error('Error updating todo:', error);
            });
    };

    return (
        <div>
            <h1>Edit Todo</h1>
            <form onSubmit={handleSubmit}>
                <label>
                    Name:
                    <input type="text" name="Name" value={todo.Name} onChange={handleChange} />
                </label>
                <label>
                    Text:
                    <input type="text" name="Text" value={todo.Text} onChange={handleChange} />
                </label>
                <label>
                    Time Limit (in days):
                    <input type="number" name="TimeLimit" value={todo.TimeLimit} onChange={handleChange} />
                </label>
                <button type="submit">Update Todo</button>
            </form>
        </div>
    );
}

export default EditTodo;
