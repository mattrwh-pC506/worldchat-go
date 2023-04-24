import React, { useEffect } from 'react';
import {isAuthorized, TOKEN_STORAGE_KEY} from "./auth";
import {useNavigate} from "react-router-dom";
import styled from "styled-components";
import {PrimaryButton} from "./components/Buttons";

const Page = styled.div`
    width: 100vw;
    height: 100vh;
    background-color: #080808;
`;

const Header = styled.div`
    width: 100vw;
    background-color: #000000;
    display: flex;
    justify-content: flex-end;
    align-items: flex-start;
`;


function App() {
    const navigate = useNavigate()
    const handleLogout = () => {
        localStorage.removeItem(TOKEN_STORAGE_KEY);
        navigate("/login");
    }

   useEffect(() => {
       isAuthorized().then((authorized) => {
           if (!authorized) {
                navigate("/login")
           }
       })
   }, [])

  return (
    <Page>
      <Header>
          <PrimaryButton onClick={handleLogout}>Logout</PrimaryButton>
      </Header>
    </Page>
  );
}

export default App;
