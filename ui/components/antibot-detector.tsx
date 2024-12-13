"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card, CardContent } from "@/components/ui/card";

export default function AntibotDetector() {
  const [url, setUrl] = useState("https://www.nike.com/");
  const [results, setResults] = useState<string[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError("");
    setResults([]);

    try {
      const response = await fetch("http://localhost:5000/api/detect", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ url }),
      });

      if (response.status !== 200) {
        throw new Error("Failed to fetch data");
      }

      const data = await response.json();
      setResults(data);
    } catch (err) {
      setError("An error occurred while fetching the data.");
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Card>
      <CardContent className="p-6">
        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="flex space-x-2">
            <Input
              type="url"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              placeholder="Enter website URL"
              required
              className="flex-grow"
            />
            <Button type="submit" disabled={loading}>
              {loading ? "Detecting..." : "Detect"}
            </Button>
          </div>
        </form>

        {error && <p className="text-red-500 mt-4">{error}</p>}

        {results.length > 0 && (
          <div className="mt-4">
            <h2 className="text-lg font-semibold mb-2">
              Detected Antibot Providers:
            </h2>
            <ul className="list-disc list-inside">
              {results.map((provider, index) => (
                <li key={index}>{provider}</li>
              ))}
            </ul>
          </div>
        )}

        {results.length === 0 && !loading && !error && (
          <p className="mt-4 text-gray-500">No antibot providers detected.</p>
        )}
      </CardContent>
    </Card>
  );
}
