import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import React from 'react';

type MyCardProps = React.ComponentPropsWithoutRef<'div'> & {
  cardTitle: string;
  cardDescription: string;
  cardContent: string;
  cardFooter: string;
  cardImage: string;
  cardImageAlt?: string;
  className?: React.ComponentPropsWithoutRef<'div'>['className'];
};

function MyCard({
  cardTitle,
  cardDescription,
  cardContent,
  cardFooter,
  cardImage,
  cardImageAlt,
  className,
  ...props
}: MyCardProps) {
  const cardImageAltText = cardImageAlt || 'Cyberpunk City';

  return (
    <>
      <Card className={className} {...props}>
        <CardHeader>
          <CardTitle>{cardTitle}</CardTitle>
          <CardDescription>{cardDescription}</CardDescription>
        </CardHeader>
        <CardContent>
          <img className='m-auto' src={cardImage} alt={cardImageAltText} />
          <p>{cardContent}</p>
        </CardContent>
        <CardFooter>
          <p>{cardFooter}</p>
        </CardFooter>
      </Card>
    </>
  );
}

export default MyCard;
