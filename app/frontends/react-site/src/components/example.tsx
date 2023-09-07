import { useAuth } from '@clerk/clerk-react';
import { Button } from '@/components/ui/button';

export default function Example() {
  const { isLoaded, userId, sessionId, getToken } = useAuth();

  const handleClick = async () => {
    const res = await fetch('https://new-service.fly.dev/test/auth', {
      headers: { Authorization: `Bearer ${await getToken()}` },
    })
      .then((res) => res.json())
      .catch((err) => console.log(err));
    console.log(res);
  };

  // In case the user signs out while on the page.
  if (!isLoaded || !userId) {
    return <h4>Loading...</h4>;
  }

  getToken().then((token) => {
    console.log(token);
  });

  return (
    <>
      <div>
        Hello, {userId} your current active session is {sessionId}
      </div>
      <Button onClick={() => getToken().then((token) => console.log(token))}>
        Get a Token
      </Button>
      <Button onClick={handleClick}>test query</Button>
    </>
  );
}
