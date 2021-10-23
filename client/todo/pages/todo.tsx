import { NextPage } from "next";
import router from "next/router";
import React, { useState } from "react";
import { postList, postTask } from "../lib/requests";

let listsCounter = -1
let tasksCounter = -1

interface Task {
    id: number;
    name: string;
    description: string;
    status: string;
}

interface List {
    id: number;
    name: string;
    tasks: Array<Task>
}


const Lists = new Array<List>()



const ToDo: NextPage = () => {
    const [listUserInput, setUserInput] = useState<string>('')
    const [todo, setTodoList] = useState<Array<List>>([])

    const [taskUserInput, setTaskUserInput] = useState<string>('')
    const [list, setList] = useState<Array<Task>>([])


    const handleChange = (e) => {
        e.preventDefault()

        setUserInput(e.target.value)
    }

    const handleSubmit = (e) => {

        e.preventDefault()

        listsCounter++
        let obj = {
            id: listsCounter,
            name: listUserInput,
            tasks: new Array<Task>(),
        }

        Lists.push(obj)

        console.log(Lists[listsCounter])

        setTodoList([...todo,
            obj
        ])

        postList(obj.name)
        // setTodoList([
        //     listUserInput,
        //     ...todo
        // ])
    }


    const handleTaskChange = (e) => {
        e.preventDefault()

        setTaskUserInput(e.target.value)
    }

    const handleTaskSubmit = (e, idx) => {
        e.preventDefault()
        tasksCounter++
        let obj = {
            id: tasksCounter,
            name: taskUserInput,
            description: "",
            status: "open",
        }
        Lists[idx].tasks.push(obj)

        postTask(obj.name, obj.description, idx)
        setList([
            ...list,
            obj,
        ])
    }

    // const handleRemove = (e, taskId, listId) => {
    //     e.preventDefault()
    //     const index =  Lists[listId].tasks.indexOf(taskId, 0);
    //     if (index > -1) {
    //         Lists[listId].tasks.splice(index, 1);
    //     }
    //     setList([
    //         ...list,
    //     ])
    // }


    return (
        <div className="h-100 w-full flex flex-col items-center justify-center bg-teal-lightest font-sans">


            <div className="shadow flex w-6/12 m-4">
                <input className="w-full rounded p-2" type="text" placeholder="Search..." />
                <button className="bg-white w-auto flex justify-end items-center text-blue-500 p-2 hover:text-blue-400">
                    <i className="material-icons">search</i>
                </button>
            </div>
            <div className="bg-white rounded shadow p-6 m-4 w-full lg:w-3/4 lg:max-w-lg">
                <div className="mb-4">
                    <h1 className="text-grey-darkest">Your lists:</h1>
                    <div className="flex mt-4">
                        <input onChange={handleChange} className="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-grey-darker" placeholder="Add List" />
                        <button onClick={handleSubmit} className="flex-no-shrink p-2 border-2 rounded text-teal border-teal bg-yellow-400 hover:bg-yellow-500 focus:bg-yellow-700">Create</button>
                    </div>
                </div>
                <div>
                    {
                        todo.map((listName, idx) => {
                            return (
                                <div key={idx} className="bg-white rounded shadow p-6 m-4 w-full lg:w-3/4 lg:max-w-lg">
                                    <div className="mb-4">
                                        <h1 className="text-grey-darkest">{listName.name}</h1>
                                        <div className="flex mt-4">
                                            <input onChange={handleTaskChange} className="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-grey-darker" placeholder="Add Task" />
                                            <button onClick={(e) => handleTaskSubmit(e, idx)} className="flex-no-shrink p-2 border-2 rounded text-teal border-teal bg-yellow-400 hover:bg-yellow-500 focus:bg-yellow-700">Add</button>
                                        </div>
                                    </div>
                                    <div>
                                        {
                                            todo[listName.id].tasks.map((taskName, id) => {
                                                return (
                                                    <div key={id} className="flex mb-4 items-center">
                                                        <p className="w-full text-grey-darkest">{taskName.name}</p>
                                                        {/* <button className="flex-no-shrink p-2 ml-4 mr-2 border-2 rounded hover:text-white text-green border-green hover:bg-green">Done</button> */}
                                                        <button className="flex-no-shrink p-2 ml-2 border-2 rounded text-red border-red bg-red-400 hover:bg-red-500  focus:bg-red-700">Remove</button>
                                                        {/* onClick={(e) => handleRemove(e, id, listName.id)} */}
                                                    </div>
                                                )
                                            })

                                        }
                                    </div>

                                </div>
                            )
                        })
                    }
                </div>
            </div>
        </div>
    )
}

export default ToDo;