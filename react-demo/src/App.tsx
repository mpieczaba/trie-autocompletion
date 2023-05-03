import React from "react";
import styled from "styled-components";

import { Input } from "./components/Input";

const StyledApp = styled.div`
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100vh;
  align-items: center;
  background-color: #212121;
`;

function App() {
  return (
    <StyledApp>
      <Input />
    </StyledApp>
  );
}

export default App;
