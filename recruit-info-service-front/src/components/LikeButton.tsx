import React, { useState, useEffect } from 'react';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHeart } from "@fortawesome/free-solid-svg-icons";
import { useLike } from '@/hooks/useLike';
import LoadingSpinner from './LoadingSpinner';

type LikeButtonProps = {
  companyId?: number;
  className?: string;
};

const LikeButton: React.FC<LikeButtonProps> = ({ companyId, className }) => {
  const [liked, setLiked] = useState<boolean | undefined>(false);
  const { createLike, deleteLike, checkLike, loading, error } = useLike();

  useEffect(() => {
    const fetchLikeStatus = async () => {
      if (companyId) {
        const isLiked = await checkLike(companyId);
        setLiked(isLiked);
      }
    };

    fetchLikeStatus();
  }, [companyId]);

  const toggleLike = async () => {
    if (!companyId || loading) return;

    try {
      if (liked) {
        await deleteLike(companyId);
      } else {
        await createLike(companyId);
      }
      setLiked(!liked);
    } catch (err) {
      switchErrorHandling(err);
    }
  };

  if (liked === null) {
    return <LoadingSpinner />;
  }

  if (!companyId) {
    return (
      <div className="text-gray-400 cursor-not-allowed">
        <FontAwesomeIcon icon={faHeart} size="2x" />
      </div>
    );
  }

  return (
    <button onClick={toggleLike} className={`like-button-class ${className}`}>
      <FontAwesomeIcon icon={faHeart} size="2x" className={liked ? "text-red-500" : "text-gray-400 hover:text-red-500"} />
    </button>
  );
};

export default LikeButton;
function switchErrorHandling(err: unknown) {
  throw new Error('Function not implemented.');
}

