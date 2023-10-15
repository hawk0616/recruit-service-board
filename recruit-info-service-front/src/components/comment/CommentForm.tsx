import { useComment } from '../../hooks/useComment';
import { useState } from 'react'

export const CommentForm = ({ companyId }: { companyId: number }) => {
  const { createComment, loading, error } = useComment();
  const [content, setContent] = useState('');
  const [showForm, setShowForm] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await createComment(companyId, content);
    setContent('');
  };

  return (
    <div className="space-y-6">
    <button 
        onClick={() => setShowForm(!showForm)}
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded my-2"
    >
        コメントする
    </button>

    {showForm && (
        <form onSubmit={handleSubmit} className="space-y-4">
            <textarea 
                value={content} 
                onChange={(e) => setContent(e.target.value)}
                className="w-full h-24 p-2 border rounded my-2 shadow-sm focus:ring focus:ring-opacity-50 focus:ring-blue-300 focus:border-blue-300"
            ></textarea>
            <button 
                type="submit" 
                disabled={loading}
                className={`w-full bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded ${loading ? 'opacity-50 cursor-not-allowed' : ''}`}
            >
                投稿
            </button>
        </form>
    )}
</div>

)

};

