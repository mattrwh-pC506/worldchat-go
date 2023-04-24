import React, { useEffect } from 'react';
import {isAuthorized, TOKEN_STORAGE_KEY} from "./auth";
import {useNavigate} from "react-router-dom";
import styled from "styled-components";
import {PrimaryButton} from "./components/Buttons";


export default function App(): JSX.Element {
    const navigate = useNavigate()

   useEffect(() => {
       isAuthorized().then((authorized) => {
           if (authorized) {
                navigate("/chat")
           } else {
               navigate("/login")
           }
       })
   }, [])

  return (
    <div />
  );
}