import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"

interface ProfileCardProps {
    title: string;
    description: string;
    content: string;
    footer: string;
    image: string;
}

function ProfileCard({ title, description, content, footer, image }: ProfileCardProps) {
    return (
        <>
            <Card>
                <CardHeader>
                    <CardTitle>{title}</CardTitle>
                    <CardDescription>{description}</CardDescription>
                </CardHeader>
                <CardContent>
                    <img className="m-auto" src={image} alt="random" />
                    <p>{content}</p>
                </CardContent>
                <CardFooter>
                    <p>{footer}</p>
                </CardFooter>
            </Card>
        </>
    )
}

export default ProfileCard