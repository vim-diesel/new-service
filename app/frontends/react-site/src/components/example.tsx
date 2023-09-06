import { useAuth } from '@clerk/clerk-react';

export default function Example() {
  const { isLoaded, userId, sessionId, getToken } = useAuth();

  // In case the user signs out while on the page.
  if (!isLoaded || !userId) {
    return null;
  }

  getToken().then((token) => {
    console.log(token);
  });

  return (
    <>
      <div>
        Hello, {userId} your current active session is {sessionId}
      </div>
    </>
  );
}
