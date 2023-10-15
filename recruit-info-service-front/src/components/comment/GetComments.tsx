import { useComment } from '@/hooks/useComment';
import React, { useState, useEffect } from 'react';
import LoadingSpinner from '../LoadingSpinner';
import { Comment } from '@/types/comment';

interface CommentsListProps {
  companyId: number;
}

const CommentsList: React.FC<CommentsListProps> = ({ companyId }) => {
  const { getCommentsByCompanyId, loading, error } = useComment();
  const [comments, setComments] = useState<Comment[]>([]);

  useEffect(() => {
    const fetchComments = async () => {
      const result = await getCommentsByCompanyId(companyId);
      if (result && result.length > 0) {
        setComments(result);
      }
    };

    fetchComments();
  }, [companyId]);

  return (
    <div className="space-y-4 p-4">
    {loading && <LoadingSpinner />}

    {comments.length > 0 ? (
        <div className="space-y-4">
            {comments.map((comment, index) => (
                <div key={index} className="p-4 bg-white rounded shadow-md">
                    <p className="text-gray-700">{comment.content}</p>
                </div>
            ))}
        </div>
    ) : (
        <p className="text-gray-500">No comments found for this company.</p>
    )}
</div>
  );
};

export default CommentsList;
