import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

function TodoList() {
    const [todos, setTodos] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/Todo/list')
            .then(response => {
                console.log(response.data);  // レスポンスの内容を確認
                setTodos(response.data.list);
            })
            .catch(error => console.error('Error loading the todos:', error));
    }, []);

    return (
        <div>
            <h1>Todo List</h1>
            {todos.length > 0 ? (
                <ul>
                    {todos.map(todo => (
                        <li key={todo.ID}>
                            {todo.Text} - <Link to={`/edit/${todo.ID}`}>Edit</Link>
                        </li>
                    ))}
                </ul>
            ) : (
                <p>No todos found.</p>
            )}
        </div>
    );

}

export default TodoList;
