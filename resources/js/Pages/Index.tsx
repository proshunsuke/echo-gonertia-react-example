import React, { FC } from 'react';
import { Head } from '@inertiajs/react';

type Props = {
    text: string
};

const Home: FC<Props> = ({text}) => {
    return (
      <>
          <Head title="Home"/>
          <h1 className="mb-1 text-xl font-bold">{text}</h1>
      </>
    );
};

export default Home;
