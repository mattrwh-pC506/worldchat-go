import Input from '@mui/material/Input';
import styled from "styled-components";
import {OutlinedInput} from "@mui/material";
import {BLACK, GREEN, LIGHT_GREEN} from "../styles";

export const InputField = styled(Input)`
    && {
      color: ${GREEN};
      background-color: transparent;
      border: none;
      border-bottom: 2px solid ${GREEN};
      font-size: 16px;
      padding: 10px;
      margin-bottom: 20px;
      width: 100%;
      outline: none;
  }
`;


export const ChatField = styled(OutlinedInput)`
    && {
      color: ${BLACK};
      background-color: ${LIGHT_GREEN};
      border-color: ${LIGHT_GREEN};
      font-size: 16px;
      width: 100%;
      border-radius: 50px;
      height: 40px;
  }
`;