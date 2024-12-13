import AntibotDetector from "@/components/antibot-detector";

export default function Home() {
  return (
    <main className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">peekr</h1>
      <AntibotDetector />
    </main>
  );
}
