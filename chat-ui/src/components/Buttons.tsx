import Button from "@mui/material/Button";
import styled from "styled-components"
import {BLACK, DISABLED_GREEN, GREEN} from "../styles";

export const PrimaryButton = styled(Button)`
    && {
      color: ${GREEN};
      border: none;
      
      &:hover:not(:disabled) {
        background-color: ${BLACK};
      }
      
      &:disabled {
        color: ${DISABLED_GREEN};
        cursor: not-allowed;
      }
  }
`;