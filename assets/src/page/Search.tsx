import { Input } from "@/components/ui/input";

function Search() {
  return (
    <div className="flex flex-col justify-center">
      <h1 className="text-4xl text-balance">Search</h1>

      <form>
        <Input type="text" />
      </form>
    </div>
  );
}

export default Search;
