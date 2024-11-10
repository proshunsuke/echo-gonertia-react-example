import React, { FC } from 'react';
import { Head } from '@inertiajs/react';

type Props = {
    status: number
};

const Error: FC<Props> = ({ status }) => {
    const title = {
        503: '503: Service Unavailable',
        500: '500: Server Error',
        404: '404: Page Not Found',
        403: '403: Forbidden',
    }[status] || 'Unknown Error';

    const description = {
        503: 'Sorry, we are doing some maintenance. Please check back soon.',
        500: 'Whoops, something went wrong on our servers.',
        404: 'Sorry, the page you are looking for could not be found.',
        403: 'Sorry, you are forbidden from accessing this page.',
    }[status] || 'An unknown error occurred.';

    return (
      <>
          <Head title="Error"/>
          <h1 className="mb-1 text-xl font-bold">{title}</h1>
          <div>{description}</div>
      </>
    );
};

export default Error;
