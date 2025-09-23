import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { ExternalLink } from "lucide-react";

interface SearchResultProps {
  containerClassName?: React.ComponentProps<"div">["className"];
  title?: string;
}

function SearchResult(props: SearchResultProps) {
  // add popover if filename is too long
  return (
    <div className={props.containerClassName}>
      <Card className="py-0 gap-0">
        <CardHeader className="border-b-1 py-3 flex flex-row justify-between content-center">
          <CardTitle className="inline-block py-1">Test</CardTitle>
          <div className="hover:bg-neutral-100 hover:p-1 hover:mr-0 hover:cursor-pointer active:bg-neutral-200 rounded-sm my-auto mr-1">
            <ExternalLink size={14} className="text-neutral-500" />
          </div>
        </CardHeader>
        <CardContent className="bg-neutral-100 py-3 rounded-b-xl">
          Hello world
        </CardContent>
      </Card>
    </div>
  );
}

export default SearchResult;
