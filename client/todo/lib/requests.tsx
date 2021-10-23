import axios from 'axios';

const apiURL = "http://127.0.0.1:8080"

declare global {
    var jwt: string;
}

export const registerUser = async (email, password) => {
    try {
        const { data } = await axios.post('http://127.0.0.1:8080/user/signup', { email, password });
        console.log(data);
        alert(data);
    } catch (err) {
        //
        alert("wrong input");
    }

}

export const loginUser = async (email, password) => {
    const { data } = await axios.post('http://127.0.0.1:8080/user/signin', { email, password });
    console.log(data);
    // authAxios = axios.create({
    //     baseURL: apiURL,
    //     headers: {
    //         Authorization: `Bearer ${data}`
    //     }
    // })
    alert(data);
    global.jwt = `${data}`
    console.log(data);
}

export const postList = async (list_name) => {
    try {
        const { data } = await axios.post(`http://127.0.0.1:8080/lists`, { list_name }, {
            headers: {
                Authorization: 'Bearer ' + jwt
            }
        });
        console.log(data);
    } catch (err) {
        //
    }
}

export const putList = async (list_name, list_id) => {
    try {
        const { data } = await axios.put(`http://127.0.0.1:8080/lists/${list_id}`, { list_name }, {
            headers: {
                Authorization: 'Bearer ' + jwt
            }
        });
        console.log(data);
    } catch (err) {
        //
    }
}

export const deleteList = async (list_id) => {
    try {
        const { data } = await axios.delete(`http://127.0.0.1:8080/lists/${list_id}`, {
            headers: {
                Authorization: 'Bearer ' + jwt
            }
        });
        console.log(data);
    } catch (err) {
        //
    }
}

export const postTask = async (task_name, task_description, list_id) => {
    try {
        const { data } = await axios.post(`http://127.0.0.1:8080/lists/${list_id}/tasks/`, {task_name, task_description}, {
            headers: {
                Authorization: 'Bearer ' + jwt
            }
        });
        console.log(data);
    } catch (err) {
        //
    }
}

export const putTask = async (list_id, task_id, task_name, status) => {
    try {
        const { data } = await axios.put(`http://127.0.0.1:8080/lists/${list_id}/tasks/${task_id}`, {task_name, status}, {
            headers: {
                Authorization: 'Bearer ' + jwt
            }
        });
        console.log(data);
    } catch (err) {
        //
    }
}

export const deleteTask= async (list_id, task_id) => {
    try {
        const { data } = await axios.delete(`http://127.0.0.1:8080/lists/${list_id}/tasks/${task_id}`, {
            headers: {
                Authorization: 'Bearer ' + jwt
            }
        });
        console.log(data);
    } catch (err) {
        //
    }
}




