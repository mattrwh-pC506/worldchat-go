import React from 'react';
import {useAuthorizedRouting} from "./router";


export default function App(): JSX.Element {
    useAuthorizedRouting();

    return (
        <div />
    );
}