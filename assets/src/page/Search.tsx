import SearchResult from "@/components/SearchResult";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useEffect, useRef, useState } from "react";

interface FormElements extends HTMLFormControlsCollection {
  searchInput: HTMLInputElement;
}
interface SearchFormElement extends HTMLFormElement {
  readonly elements: FormElements;
}

function Search() {
  const [hasSearched, setHasSearched] = useState<boolean>(false);
  const searchInputRef = useRef<HTMLInputElement | null>(null);

  useEffect(() => {
    const handleShortcuts = (event: KeyboardEvent) => {
      if (event.ctrlKey && event.key === "k") {
        event.preventDefault();
        if (searchInputRef.current) {
          searchInputRef.current.focus();
          searchInputRef.current.value = "";
        }
      }
    };

    window.addEventListener("keydown", handleShortcuts);

    return () => {
      window.removeEventListener("keydown", handleShortcuts);
    };
  });

  const handleSubmit = async (e: React.FormEvent<SearchFormElement>) => {
    e.preventDefault();
    const searchValue = e.currentTarget.elements.searchInput.value;
    if (searchValue !== "") {
      setHasSearched(true);

      await fetch("/api/search", {
        method: "POST",
        body: JSON.stringify({
          query: searchValue,
        }),
      });
    }
  };

  return (
    <div className="flex flex-col justify-start w-full">
      <div
        className={`w-3xl ${hasSearched ? "justify-start mt-5 mx-auto" : "mt-120 mx-auto"}`}
      >
        <form onSubmit={handleSubmit} className="flex gap-2">
          <Input
            ref={searchInputRef}
            name="searchInput"
            type="text"
            placeholder="Search your code..."
          />
          <Button variant="outline" type="submit">
            Search
          </Button>
        </form>
      </div>

      {/*todo*/}
      {/* <Skeleton className="h-100 w-full rounded-xl mt-10" /> */}

      {hasSearched && (
        <div className="size-auto mx-10 mt-10 flex flex-col gap-6">
          <SearchResult title="test" content="test" />
          <SearchResult title="test" content="test" />
          <SearchResult title="test" content="test" />
          <SearchResult title="test" content="test" />
        </div>
      )}
    </div>
  );
}

export default Search;
