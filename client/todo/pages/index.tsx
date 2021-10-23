import { NextPage } from 'next/types';
import Link from 'next/link';
import React from 'react';
import { registerUser } from '../lib/requests';
import Router, { useRouter } from 'next/router';



const Register: NextPage = () => {
    const router = useRouter()

    const email = React.useRef<HTMLInputElement>(null);
    const password = React.useRef<HTMLInputElement>(null);

    const [success, setSuccess] = React.useState<boolean>(false);


    const handleSubmit = (e) => {

        let err: boolean = false;

        const inputEmail = email.current?.value || '';
        const inputPassword = password.current?.value || '';

        if (!inputEmail) {
            alert("Please input your email address");
            err = true;
        } else if (inputPassword === "") {
            alert("Please provide your password!");
            err = true;
        } else if (inputPassword.length <= 7) {
            alert("Your password must have 8 characters or greater");
            err = true;
        } else {
            registerUser(inputEmail, inputPassword)
            setSuccess(true);
            e.preventDefault()
            router.push('/login')
        }


    };

    // React.useEffect(() => {
    //     if (success) {
    //         Router.push('/login');
    //     }
    // }, []);

    return (
        <div className="bg-openware-back bg-cover flex items-center justify-center h-screen">
            <div className="absolute bg-cover h-screen w-full bg-dino fill-current text-black z-0"/>
            <div className="w-full max-w-md z-10">
                <form onSubmit={handleSubmit} className="bg-white shadow-lg rounded px-12 pt-6 pb-8 mb-4">
                    <div
                        className="text-gray-800 text-2xl flex justify-center border-b-2 py-2 mb-4"
                    >
                        Register
                    </div>
                    <div className="mb-4">
                        <label
                            className="block text-gray-700 text-sm font-normal mb-2"
                            htmlFor="email"
                        >
                            Email
                        </label>
                        <input
                            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            ref={email}
                            name="email"
                            type="email"
                            placeholder="Email"
                        />
                    </div>
                    <div className="mb-6">
                        <label
                            className="block text-gray-700 text-sm font-normal mb-2"
                            htmlFor="password"
                        >
                            Password
                        </label>
                        <input
                            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
                            type="password"
                            ref={password}
                            placeholder="Password"
                            name="password"
                        />
                    </div>
                    <div className="flex items-center justify-between">
                        <button className="px-4 py-2 rounded text-black inline-block shadow-lg bg-yellow-400 hover:bg-yellow-500 focus:bg-yellow-700" type="submit">Sign Up</button>
                    </div>
                </form>
            </div>
        </div>

    )

}


export default Register;