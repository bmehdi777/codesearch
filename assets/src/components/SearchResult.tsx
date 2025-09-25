import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { ExternalLink } from "lucide-react";
import { useMemo } from "react";
import { toast } from "sonner";

interface SearchResultProps {
  containerClassName?: React.ComponentProps<"div">["className"];
  title: string;
  content: string;
}

function SearchResult(props: SearchResultProps) {
  const splitContent = useMemo(
    () => props.content.split("\n"),
    [props.content],
  );

  const externalClick = () => {
    toast(`Opened nvim with \`${props.title}\``);

		// call backend to open nvim with this file
  };

  return (
    <Card className="py-0 gap-0 w-full">
      <CardHeader className="border-b-1 py-3 flex flex-row justify-between content-center">
        <CardTitle className="inline-block py-1">{props.title}</CardTitle>
        <div className="hover:bg-neutral-100 hover:p-1 hover:mr-0 hover:cursor-pointer active:bg-neutral-200 rounded-sm my-auto mr-1">
          <ExternalLink
            size={14}
            className="text-neutral-500"
            onClick={externalClick}
          />
        </div>
      </CardHeader>
      <CardContent className="bg-neutral-100 py-3 rounded-b-xl">
        {splitContent.map((content) => (
          <pre className="w-full truncate">{content}</pre>
        ))}
      </CardContent>
    </Card>
  );
}

export default SearchResult;
